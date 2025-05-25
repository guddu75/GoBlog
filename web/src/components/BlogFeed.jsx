import React from 'react';
import BlogPost from './BlogPost';


const BlogFeed = () => {
  const blogPosts = [
    {
      id: 1,
      title: 'Getting Started with React',
      author: 'Alice Johnson',
      date: 'May 20, 2025',
      content: (
        <>
          <p>React is a powerful JavaScript library for building user interfaces. It allows developers to create large web applications that can change data without reloading the page. The main advantage of React is its use of a virtual DOM, which makes updates very efficient.</p>
          <p>Key concepts include components, props, state, and hooks. Components are reusable UI pieces, props are used to pass data down to components, and state is used to manage data that changes over time within a component.</p>
          <p>This blog post is a simple example of how to display content dynamically using React components.</p>
        </>
      ),
      imageUrl: 'https://placehold.co/600x400/F4F4F4/333333?text=React+Basics',
      imageAlt: 'Abstract image representing React development'
    },
    {
      id: 2,
      title: 'The Power of Tailwind CSS',
      author: 'Bob Williams',
      date: 'May 22, 2025',
      content: (
        <>
          <p>Tailwind CSS is a utility-first CSS framework for rapidly building custom user interfaces. Unlike traditional CSS frameworks that come with pre-designed components, Tailwind provides low-level utility classes that you can combine to build any design directly in your markup.</p>
          <p>This approach leads to highly customizable designs and often faster development cycles, as you spend less time writing custom CSS and more time composing existing utilities. It's especially popular in React applications for its flexibility.</p>
          <p>For example, to center text, you can just add `text-center` class. To add padding, `p-4`. It's incredibly intuitive once you get the hang of it.</p>
        </>
      ),
      imageUrl: 'https://placehold.co/600x400/E0F2F7/007BFF?text=Tailwind+CSS',
      imageAlt: 'Abstract image representing Tailwind CSS'
    },
    {
      id: 3,
      title: 'Understanding JSX in React',
      author: 'Charlie Brown',
      date: 'May 24, 2025',
      content: (
        <>
          <p>JSX stands for JavaScript XML. It's a syntax extension for JavaScript, recommended by React, that allows you to write HTML-like markup directly within your JavaScript code. This makes it easier to visualize the UI structure directly alongside the logic that renders it.</p>
          <p>While JSX looks like HTML, it's not. It gets transpiled into regular JavaScript function calls (e.g., `React.createElement`). This means you can embed JavaScript expressions within your JSX using curly braces `{}`.</p>
          <p>For instance, you use `className` instead of `class` for CSS classes, and `htmlFor` instead of `for` for labels, because `class` and `for` are reserved keywords in JavaScript.</p>
        </>
      ),
      imageUrl: 'https://placehold.co/600x400/D1E7DD/28A745?text=JSX+Syntax',
      imageAlt: 'Abstract image representing JSX syntax'
    },
    {
      id: 4,
      title: 'State Management with Hooks',
      author: 'David Lee',
      date: 'May 26, 2025',
      content: (
        <>
          <p>Managing state in React applications is crucial. React Hooks, introduced in React 16.8, provide a way to use state and other React features without writing a class. The `useState` hook allows you to add state to functional components.</p>
          <p>For side effects like data fetching or subscriptions, the `useEffect` hook is your go-to. It runs after every render, but you can control when it re-runs by providing a dependency array.</p>
          <p>Hooks simplify component logic and make it more readable and reusable.</p>
        </>
      ),
      imageUrl: 'https://placehold.co/600x400/D7E9F7/1A5276?text=React+Hooks',
      imageAlt: 'Abstract image representing React Hooks'
    },
    {
      id: 5,
      title: 'Building Responsive Layouts',
      author: 'Eve Green',
      date: 'May 28, 2025',
      content: (
        <>
          <p>In today's multi-device world, responsive web design is not just a luxury, but a necessity. Building layouts that adapt seamlessly to different screen sizes ensures a great user experience on desktops, tablets, and mobile phones.</p>
          <p>Tailwind CSS makes responsive design incredibly easy with its utility-first approach and responsive prefixes (e.g., `sm:`, `md:`, `lg:`). You can define different styles for various breakpoints directly in your HTML.</p>
          <p>Techniques like Flexbox and CSS Grid are fundamental for creating flexible and dynamic layouts that respond gracefully to changes in viewport size.</p>
        </>
      ),
      imageUrl: 'https://placehold.co/600x400/F0F8FF/4682B4?text=Responsive+Design',
      imageAlt: 'Abstract image representing responsive web design'
    },
  ];

  return (
    <div className="container mx-auto p-4 md:p-8">
      <h2 className="text-4xl md:text-5xl font-extrabold text-green-800 text-center mb-10">
        Latest Blog Posts
      </h2>
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {blogPosts.map((post) => (
          <BlogPost
            key={post.id}
            title={post.title}
            author={post.author}
            date={post.date}
            content={post.content}
            imageUrl={post.imageUrl}
            imageAlt={post.imageAlt}
          />
        ))}
      </div>
    </div>
  );
};

export default BlogFeed;
