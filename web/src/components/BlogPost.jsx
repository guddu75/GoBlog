import React from 'react';

const BlogPost = ({ title, author, date, content, imageUrl, imageAlt }) => {

    content = content.substr(200)+"..." || "No content available"; 
    console.log(content);// Fallback for content if not provided

  return (
    // Article container with responsive styling for a blog post card
    <article className="bg-white rounded-xl shadow-lg hover:shadow-xl transition-shadow duration-300 overflow-hidden flex flex-col">
      {/* Blog post image, conditionally rendered if imageUrl is provided */}
      {imageUrl && (
        <img
          src={imageUrl}
          alt={imageAlt || title} // Use imageAlt if provided, otherwise fallback to title
          className="w-full h-48 object-cover" // Responsive image styling
          // Fallback for image loading errors: displays a placeholder if the image fails to load
          onError={(e) => {
            e.target.onerror = null; // Prevents infinite loop if placeholder also fails
            e.target.src = `https://placehold.co/600x400/CCCCCC/333333?text=Image+Error`; // Generic placeholder
          }}
        />
      )}

      {/* Content section of the blog post card */}
      <div className="p-6 flex flex-col flex-grow">
        {/* Title of the blog post */}
        <h2 className="text-2xl font-bold text-gray-800 mb-2">{title}</h2>
        {/* Metadata: Author and Date */}
        <p className="text-sm text-gray-500 mb-4">
          By <span className="font-semibold text-blue-600">{author}</span> on {date}
        </p>
        {/* Main content of the blog post */}
        <div className="text-gray-700 leading-relaxed flex-grow">
          {content} {/* Renders the content, which can be JSX or a string */}
        </div>
        {/* Read more button */}
        <button className="mt-6 self-start bg-blue-500 hover:bg-blue-600 text-white font-semibold py-2 px-4 rounded-lg shadow-md transition-colors duration-300">
          Read More
        </button>
      </div>
    </article>
  );
};

export default BlogPost;
