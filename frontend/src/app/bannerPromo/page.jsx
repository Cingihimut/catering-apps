"use client"
import { ArrowCircleLeft, ArrowCircleRight } from '@phosphor-icons/react';
import Image from 'next/image';
import { useState } from 'react';

const Carousel = () => {
  const images = [
    '/public/assets/Banner-1.png',
    '/public/assets/Banner-2.png',
    '/public/assets/Banner-3.png',
    '/public/assets/Banner-4.png',
  ];

  const [currentImage, setCurrentImage] = useState(0);

  const nextImage = () => {
    setCurrentImage((prev) => (prev + 1) % images.length);
  };

  const prevImage = () => {
    setCurrentImage((prev) => (prev - 1 + images.length) % images.length);
  };

  return (
    <div className="relative p-4">
      <Image src={`/images/${images[currentImage]}`} alt={`Image ${currentImage + 1}`} width={700} height={300} className="w-full h-64 object-cover rounded-l" />
      <div className="absolute top-0 left-0 right-0 bottom-0 flex items-center justify-between p-4">
        <button onClick={prevImage} className="text-white bg-gray-800 px-4 py-2 rounded-full opacity-75 hover:opacity-100">
            <ArrowCircleLeft size={20} />
        </button>
        <button onClick={nextImage} className="text-white bg-gray-800 px-4 py-2 rounded-full opacity-75 hover:opacity-100">
            <ArrowCircleRight size={20} />
        </button>
      </div>
    </div>
  );
};

export default Carousel;
