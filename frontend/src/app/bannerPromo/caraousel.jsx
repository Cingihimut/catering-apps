"use client";
import { ArrowCircleLeft, ArrowCircleRight } from "@phosphor-icons/react";
import Image from "next/image";
import { useState } from "react";

const Carousel = () => {
  const images = [
    "/assets/Banner-1-min.png",
    "/assets/Banner-2-min.png",
    "/assets/Banner-3-min.png",
    "/assets/Banner-4-min.png",
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
      <Image
        src={images[currentImage]}
        alt={`Image ${currentImage + 1}`}
        width={300}
        height={100}
        className="w-full h-64 object-cover rounded-lg"
        onError={(e) => console.error(`Error loading image: ${e.target.src}`)}
      />

      <div className="absolute top-0 left-0 right-0 bottom-0 flex items-center justify-between p-4">
        <button
          onClick={prevImage}
          className="text-white bg-gray-800 px-4 py-2 rounded-full opacity-75 hover:opacity-100"
        >
          <ArrowCircleLeft size={20} />
        </button>
        <button
          onClick={nextImage}
          className="text-white bg-gray-800 px-4 py-2 rounded-full opacity-75 hover:opacity-100"
        >
          <ArrowCircleRight size={20} />
        </button>
      </div>
    </div>
  );
};

export default Carousel;
