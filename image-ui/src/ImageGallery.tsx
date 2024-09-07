import React, { useEffect, useState } from 'react';
import axios from 'axios';

interface Image {
  id: number;
  filename: string;
}

const ImageGallery: React.FC = () => {
  const [images, setImages] = useState<Image[]>([]);

  useEffect(() => {
    const fetchImages = async () => {
      try {
        const response = await axios.get<Image[]>('http://localhost:3000/images');
        // กรองรูปภาพที่ต้องการแสดง
        const filteredImages = response.data.filter(image => image.id < 5); // ตัวอย่างการกรอง
        setImages(filteredImages);
      } catch (error) {
        console.error('Error fetching images:', error);
      }
    };

    fetchImages();
  }, []);

  return (
    <div>
      <h2>Image Gallery</h2>
      <div style={{ display: 'flex', flexWrap: 'wrap' }}>
        {images.map((image) => (
          <div key={image.id} style={{ margin: '10px' }}>
            <img src={image.filename} alt={`Image ${image.id}`} style={{ width: '200px' }} />
          </div>
        ))}
      </div>
    </div>
  );
};

export default ImageGallery;
