"use client"
import React, { useState } from "react";
import Link from "next/link";

export default function Navbar() {
  const [cartCount, setCartCount] = useState(0);

  const handleAddToCart = () => {
    setCartCount(cartCount + 1);
  };

  return (
    <nav className="flex justify-between items-center p-4  ">
      <div className="flex gap-6 font-semibold text-gray-500">
        <Link href="/">
          <img className="w-30 h-16" src="/logo-navbar.png"/>
        </Link>
        <div className="flex space-x-4 my-auto">
          <Link href="/about">About</Link>
          <Link href="/contact">Contact</Link>
          <Link href="/menu">Menu</Link>
        </div>
      </div>

      <div className="flex gap-2">
        <div className="my-auto">
          <Link href={'/transaksi'}>
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            strokeWidth={1.5}
            stroke="currentColor"
            className="w-6 h-6"
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              d="M2.25 3h1.386c.51 0 .955.343 1.087.835l.383 1.437M7.5 14.25a3 3 0 00-3 3h15.75m-12.75-3h11.218c1.121-2.3 2.1-4.684 2.924-7.138a60.114 60.114 0 00-16.536-1.84M7.5 14.25L5.106 5.272M6 20.25a.75.75 0 11-1.5 0 .75.75 0 011.5 0zm12.75 0a.75.75 0 11-1.5 0 .75.75 0 011.5 0z"
            />
          </svg>
          </Link>
        </div>
        <span>{cartCount}</span>
        <Link href="/login">
          <button className=" px-4 py-2 rounded">Login</button>
        </Link>
        <Link href="/register">
          <button className="bg-color text-white px-4 py-2 rounded-full">
            Register
          </button>
        </Link>
      </div>
    </nav>
  );
}
