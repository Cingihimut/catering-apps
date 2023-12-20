// DaftarMenu.jsx
import { useState } from "react";

export default function DaftarMenu({ product }) {
  const [cartCount, setCartCount] = useState(0);

  const handleAddToCart = () => {
    setCartCount(cartCount + 1);
    alert(`Item ${product.product_name} added to cart! Price: ${product.price}`);
  };

  return (
    <div className="flex space-x-4 gap-2">
      <div
        key={product.name}
        className="bg-gray-200 p-4 rounded-lg text-center hover:bg-white transition-all duration-500 hover:shadow-2xl hover:shadow-black/30"
      >
        <div className="text-center"></div>
        <h4 className="font-semibold my-3">{product.product_name}</h4>
        <p className="text-gray-500 text-sm">{product.description}</p>
        <button
          className="bg-green-500 text-white px-4 py-2 rounded hover:bg-green-600 transition-all duration-300"
          onClick={handleAddToCart}
        >
          Add to Cart ${product.price}
        </button>
      </div>
    </div>
  );
}
