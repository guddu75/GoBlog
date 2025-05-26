import React, { useState } from 'react';
import { redirect , useNavigate } from 'react-router-dom';
import { API_URL } from '../util';

// CreatePost component using an ES6 function expression
const CreatePost = () => {
    const redirect = useNavigate(); // Hook for programmatic navigation
  const [title, setTitle] = useState('');
  const [content, setContent] = useState('');
  // Removed the 'author' state as requested
  const [tags, setTags] = useState(''); // New state for tags

  const handleSubmit = (e) => {
    e.preventDefault();
    
    // In a real application, you might want to process tags (e.g., split by comma into an array)
    const processedTags = tags.split(',').map(tag => tag.trim()).filter(tag => tag !== '');

    console.log('New Post:', {
      title,
      content,
      // Removed 'author' from console.log
      tags: processedTags // Send processed tags
    });


    const post = {
        title,
        content,
        tags: processedTags
    }
    
    const token = localStorage.getItem('token');
    if (!token) {
      alert('You need to be logged in to create a post.');
      redirect('/login'); 
      return;
    }
    
    fetch(`${API_URL}/posts`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`
      },
      body: JSON.stringify(post) 
    })
    .then(response => response.json())
    .then(data => {
      console.log('Post created successfully:', data);
      alert('Post created!');
    })
    .catch(error => {
      console.error('Error creating post:', error);
      alert('Failed to create post.');
    });
    
    setTitle('');
    setContent('');
    setTags(''); 
  };

  return (
    // Changed background back to bg-gray-100 and removed max-w-2xl from this container
    <div className="flex items-center justify-center min-h-[calc(100vh-160px)] p-4 bg-gray-100">
      {/* Removed max-w-2xl from this div to allow the form to expand */}
      <div className="bg-white p-8 rounded-xl shadow-lg w-full"> {/* Changed max-w-2xl to w-full */}
        <h2 className="text-3xl font-bold text-center text-gray-800 mb-6">Create New Blog Post</h2>
        <form onSubmit={handleSubmit} className="space-y-6">
          <div>
            <label htmlFor="title" className="block text-gray-700 text-sm font-medium mb-2">
              Post Title
            </label>
            <input
              type="text"
              id="title"
              className="text-black w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              placeholder="Enter your post title"
              value={title}
              onChange={(e) => setTitle(e.target.value)}
              required
            />
          </div>
          {/* Removed the entire author input section */}
          <div>
            <label htmlFor="content" className="block text-gray-700 text-sm font-medium mb-2">
              Post Content
            </label>
            <textarea
              id="content"
              rows="10"
              className="text-black w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              placeholder="Write your blog post here..."
              value={content}
              onChange={(e) => setContent(e.target.value)}
              required
            ></textarea>
          </div>
          <div>
            <label htmlFor="tags" className="block text-gray-700 text-sm font-medium mb-2">
              Tags (comma-separated)
            </label>
            <input
              type="text"
              id="tags"
              className="text-black w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              placeholder="e.g., react, javascript, frontend"
              value={tags}
              onChange={(e) => setTags(e.target.value)}
            />
          </div>
          <button
            type="submit"
            className="w-full bg-blue-600 hover:bg-blue-700 text-white font-bold py-3 px-4 rounded-lg transition-colors duration-300 shadow-md"
          >
            Publish Post
          </button>
        </form>
      </div>
    </div>
  );
};

export default CreatePost;
