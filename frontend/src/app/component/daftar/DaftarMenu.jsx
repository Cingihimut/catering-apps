"use client";
import { useState } from "react";

export default function DaftarMenu({ name }) {
  const [cartCount, setCartCount] = useState(0);
  const handleAddToCart = () => {
    setCartCount(cartCount + 1);
    alert(`Item ${name} added to cart!`);
  };
  return (
    <div
      className="bg-gray-200 p-4 rounded-lg text-center hover:bg-white transition-all hover:shadow-2xl hover:shadow-black/30"
      key={name}
    >
      <div className="text-center">
        <img
          className="rounded-full max-h-[150px] items-center block mx-auto"
          src={`gambar${name}.png`}
          alt="nasi goreng"
        />
      </div>
      <h4 className="font-semibold my-3">Catering {name}</h4>
      <p className="text-gray-500 text-sm">
        Lorem Ipsum is simply dummy text of the printing and typesetting
        industry.
      </p>
      <button
        className=" bg-color text-white rounded-full px-6 py-2 mt-4"
        onClick={handleAddToCart}
      >
        Add to card $1
      </button>
    </div>
  );
}
