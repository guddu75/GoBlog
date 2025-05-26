import { StrictMode } from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'
import App from './App.jsx'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import { ConfirmationPage } from './ConfirmationPage.jsx'
import Login from './components/Login.jsx'
import SignUp from './components/Signup.jsx'
import BlogFeed from './components/BlogFeed.jsx'
import CreatePost from './components/CreatePost.jsx'


const router = createBrowserRouter([
  {
    path: "/",
    element: <App />
  },
  {
    path : "/login",
    element: <Login />
  },
  {
    path : "/signup",
    element: <SignUp />
  },
  {
    path : "/feed",
    element: <BlogFeed />
  },
  {
    path: "/confirm/:token",
    element: <ConfirmationPage />
  },
  {
    path: "/create-post",
    element: <CreatePost/>
  }
]);


const root = ReactDOM.createRoot(document.getElementById('root'));

root.render(
  <StrictMode>
    <RouterProvider router={router} />
  </StrictMode>,
);
