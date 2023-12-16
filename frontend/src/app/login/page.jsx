"use client";
import React, { useState } from "react";
import "../globals.css";
import Link from "next/link";
import { useRouter } from "next/navigation";
import decodeToken from "@/utils/decodeToken";
import useStore from "@/stores/store";

const Login = () => {
  const setUser = useStore((state) => state.setUser);

  const router = useRouter();
  const [formData, setFormData] = useState({
    email: "",
    password: "",
  });
  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await fetch("http://localhost:8080/api/sellers/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(formData),
      });
      if (response.ok) {
        const data = await response.json();
        localStorage.setItem("accessToken", data.token);
        console.log(data.token);
        const decodedUser = decodeToken(data.token);
        setUser({ email: decodedUser.email, id: decodedUser.id });
        router.push("/");
      } else {
        const gagal = await response.json();
        console.log("Data respons JSON:", gagal);
        console.error("Registrasi gagal");
      }
    } catch (error) {
      console.error("Terjadi kesalahan:", error);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-r from-green-400 via-green-500 to-green-600">
      <div className="bg-white p-8 rounded-lg shadow-md ">
        <h1 className="text-3xl font-bold mb-6">Login</h1>
        <form onSubmit={handleSubmit} className="space-y-4">
          <label className="block">
            Email:
            <input
              type="email"
              name="email"
              value={formData.email}
              onChange={handleChange}
              required
              className="block w-full rounded-md border-2 border-gray-300 py-2 px-3 placeholder-gray-500 focus:outline-none focus:border-indigo-500 focus:ring focus:ring-indigo-200"
            />
          </label>
          <label className="block">
            Password:
            <input
              type="password"
              name="password"
              value={formData.password}
              onChange={handleChange}
              required
              className="block w-full rounded-md border-2 border-gray-300 py-2 px-3 placeholder-gray-500 focus:outline-none focus:border-indigo-500 focus:ring focus:ring-indigo-200"
            />
          </label>
          <p className="text-sm text-gray-600">
            Dont have an account?{" "}
            <Link href="/register" className="text-blue-500 hover:underline">
              Register here
            </Link>
          </p>
          <button
            type="submit"
            className="bg-indigo-500 text-white py-2 px-4 rounded-md hover:bg-indigo-600 focus:outline-none focus:ring focus:ring-indigo-200"
          >
            Login
          </button>
        </form>
      </div>
    </div>
  );
};
export default Login;