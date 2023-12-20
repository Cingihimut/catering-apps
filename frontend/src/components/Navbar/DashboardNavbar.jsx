"use client";
import React, { useState } from "react";
import { useUserStore } from "@/stores/userStore";
import { useRouter } from "next/navigation";
const DashboardNavbar = () => {
  const router = useRouter();
  const { setUser } = useUserStore();
  const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false);

  const toggleMobileMenu = () => {
    setIsMobileMenuOpen(!isMobileMenuOpen);
  };
  const handleLogout = () => {
    setUser(null);
    router.push("/");
  };

  return (
    <nav className="bg-primary p-4 px-24 w-screen sticky top-0 z-10">
      <div className=" mx-auto flex items-center justify-between">
        <div className="div">
          <a href="#" className="text-white text-lg font-semibold">
            Owner Dashboard
          </a>
        </div>

        <div>
          <div className="hidden md:flex space-x-6">
            <a href="#" className="text-white hover:text-gray-300">
              Produk
            </a>
            <a href="#" className="text-white hover:text-gray-300">
              Kategori
            </a>
            <a href="#" className="text-white hover:text-gray-300">
              Transaksi
            </a>
          </div>
        </div>
        <div>
          {" "}
          <button
            onClick={handleLogout}
            className="text-white hover:text-gray-300"
          >
            Logout
          </button>
        </div>

        <div className="md:hidden relative">
          <button
            onClick={toggleMobileMenu}
            className="text-white focus:outline-none"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
              className="h-6 w-6"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth="2"
                d="M4 6h16M4 12h16m-7 6h7"
              ></path>
            </svg>
          </button>

          {/* Mobile Menu (Hidden by default) */}
          <div
            className={`absolute top-full left-0 right-0 ${
              isMobileMenuOpen ? "" : "hidden"
            } bg-gray-800 p-4 space-y-4`}
          >
            <a href="#" className="text-white hover:text-gray-300">
              Produk
            </a>
            <a href="#" className="text-white hover:text-gray-300">
              Kategori
            </a>
            <a href="#" className="text-white hover:text-gray-300">
              Transaksi
            </a>
            <a href="#" className="text-white hover:text-gray-300">
              Logout
            </a>
          </div>
        </div>
      </div>
    </nav>
  );
};

export default DashboardNavbar;
