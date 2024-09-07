import React from 'react';
import { BrowserRouter as Router, Route, Link, Routes } from 'react-router-dom';
import ImageUpload from './ImageUpload';
import ImageGallery from './ImageGallery';

const App: React.FC = () => {
  return (
    <Router>
      <nav>
        <ul>
          <li>
            <Link to="/upload">Upload Image</Link>
          </li>
          <li>
            <Link to="/gallery">View Gallery</Link>
          </li>
        </ul>
      </nav>
      <Routes>
        <Route path="/upload" element={<ImageUpload />} />
        <Route path="/gallery" element={<ImageGallery />} />
      </Routes>
    </Router>
  );
};

export default App;
