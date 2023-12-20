// REGISTER
"use client";
import Image from "next/image";
import { useState } from "react";
import Link from "next/link";

const API_URL = process.env.API_URL + "/api/users/register";

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
    }
  }
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