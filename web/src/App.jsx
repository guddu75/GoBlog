import './App.css'
import Landing from './components/Landing';

export const API_URL = import.meta.env.VITE_API_URL || "http://localhost:8080/v1"

function App() {
  return (
    <>
      <Landing />
      <footer className="bg-gray-800 text-white text-center py-4 mt-8">
        <p className="text-sm">© {new Date().getFullYear()} goBlog. All rights reserved.</p>
        <p className="text-xs">Built with ❤️ by the goBlog Team</p>
      </footer>
    </>
  )
}

export default App;