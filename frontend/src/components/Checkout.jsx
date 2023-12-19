import { product } from "@/libs/product";
import React, { useState } from "react";

const Checkout = () => {
  const [cartCount, setCartCount] = useState(0);

  const handleAddToCart = () => {
    setCartCount(cartCount + 1);
    alert(`Item ${name} added to cart!`);
  };

  const checkout = async () => {
    const data = {
      id: product.id,
      productName: product.name,
      price: product.price,
      quantity: 3
    }

    const response = await fetch("/api/tokenizer", {
      method: "POST",
      body: JSON.stringify(data)
    })
    const requestData = await response.json()
    window.snap.pay(requestData.token)
    // console.log(requestData);

  };

  return (
    <>
      <button
        className=" bg-color bg-color-hover transition-all duration-300 text-white rounded-full px-6 py-2 mt-4"
        onClick={checkout}
      >
        Add to card $1
      </button>
    </>
  );
};

export default Checkout;