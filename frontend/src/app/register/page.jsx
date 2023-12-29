// REGISTER
"use client";
import React, { useState } from "react";
<<<<<<< HEAD
import "../globals.css";
import Link from "next/link";
import { useRouter } from "next/navigation";
=======
import { useRouter } from "next/router";
import "../globals.css";
>>>>>>> ab68c8cd872d067cf52898ba1c5c1ec294febc57

const Register = () => {
  const router = useRouter();
  const [formData, setFormData] = useState({
    username: "",
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
    }
  };

  return (
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
  );
}