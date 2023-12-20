// REGISTER
"use client";
<<<<<<< HEAD
import Image from "next/image";
import { useState } from "react";
import Link from "next/link";
=======
import React, { useState } from "react";
<<<<<<< HEAD
import "../globals.css";
import Link from "next/link";
import { useRouter } from "next/navigation";
=======
import { useRouter } from "next/router";
import "../globals.css";
>>>>>>> ab68c8cd872d067cf52898ba1c5c1ec294febc57
>>>>>>> main

const API_URL = process.env.API_URL + "/api/users/register";

<<<<<<< HEAD
export default function RegisterPage() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [name, setName] = useState("");
  const [creatingUser, setCreatingUser] = useState(false);
  const [userCreated, setUserCreated] = useState(false);
  async function handleFormSubmit(ev) {
    ev.preventDefault();
    setCreatingUser(true);
    const response = await fetch(API_URL, {
      method: "POST",
      body: JSON.stringify({ name, email, password }),
      headers: { "Content-Type": "application/json" },
    });
    console.log(process.env);
    if (response.ok) {
      setCreatingUser(false);
      setUserCreated(true);
=======
  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await fetch(
        "http://localhost:8080/api/sellers/register",
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(formData),
        }
      );
      if (response.ok) {
        const data = await response.json();
        console.log("Data respons JSON:", data);
<<<<<<< HEAD
        router.push("/login");
        console.log(data);
      } else {
        const gagal = await response.json();
        console.log("Data respons JSON:", gagal);
=======
        console.log(data);
        console.log("Registrasi berhasil!");
        router.push("/login");
      } else {
        const gagal = await response.json();
        console.log("Data respons JSON:", gagal);
        console.error("Registrasi gagal");
>>>>>>> ab68c8cd872d067cf52898ba1c5c1ec294febc57
      }
    } catch (error) {
      console.error("Terjadi kesalahan:", error);
>>>>>>> main
    }
  }
  return (
<<<<<<< HEAD
    <section className="mt-8">
      <h1 className="text-center color text-4xl">Register</h1>
      {userCreated && (
        <div className="my-4 text-center">
          Register berhasil. masuk ke{" "}
          <Link href={"/login"}>
            {" "}
            <span className="spanred">Login &raquo;</span>
          </Link>
        </div>
      )}
      <form className=" max-w-xs mx-auto" onSubmit={handleFormSubmit}>
        <input
          type="name"
          placeholder="Full Name"
          value={name}
          disabled={creatingUser}
          onChange={(ev) => setName(ev.target.value)}
        />
        <input
          type="email"
          placeholder="email"
          value={email}
          disabled={creatingUser}
          onChange={(ev) => setEmail(ev.target.value)}
        />
        <input
          type="password"
          placeholder="password"
          value={password}
          disabled={creatingUser}
          onChange={(ev) => setPassword(ev.target.value)}
        />
        <button
          className="bg-color text-white"
          type="submit"
          disabled={creatingUser}
        >
          Register
        </button>
        <div className="my-4 text-center text-gray-500">Or Login With</div>
        <button className="border border-solid border-2 flex gap-4 justify-center font-semibold">
          <Image src={"/google.png"} width={24} height={24} alt={"google"} />
          Login with google
        </button>
        <div className="text-center my-4 text-gray-500">
          Sudah punya akun?{" "}
          <Link className="underline" href={"/login"}>
            Login here
          </Link>
        </div>
      </form>
    </section>
=======
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-r from-green-400 via-green-500 to-green-600">
      <div className="bg-white p-8 rounded-lg shadow-md ">
        <h1 className="text-3xl font-bold mb-6">Register</h1>
        <form onSubmit={handleSubmit} className="space-y-4">
          <label className="block">
            Username:
            <input
              type="text"
              name="username"
              value={formData.username}
              onChange={handleChange}
              required
              className="block w-full rounded-md border-2 border-gray-300 py-2 px-3 placeholder-gray-500 focus:outline-none focus:border-indigo-500 focus:ring focus:ring-indigo-200"
            />
          </label>
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
<<<<<<< HEAD
            Already have an account?
            <Link href="/login" className="text-indigo-500 hover:underline">
              Login here
            </Link>
=======
            Already have an account?{" "}
            <a href="/login" className="text-indigo-500 hover:underline">
              Login here
            </a>
>>>>>>> ab68c8cd872d067cf52898ba1c5c1ec294febc57
          </p>
          <button
            type="submit"
            className="bg-indigo-500 text-white py-2 px-4 rounded-md hover:bg-indigo-600 focus:outline-none focus:ring focus:ring-indigo-200"
          >
            Register
          </button>
        </form>
      </div>
    </div>
>>>>>>> main
  );
}
