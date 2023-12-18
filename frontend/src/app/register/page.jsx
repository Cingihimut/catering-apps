// REGISTER
"use client";
import Image from "next/image";
import { useState } from "react";
import Link from "next/link";

export default function RegisterPage() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [creatingUser, setCreatingUser] = useState(false);
  const [userCreated, setUserCreated] = useState(false);
  async function handleFormSubmit(ev) {
    ev.preventDefault();
    setCreatingUser(true);
    await fetch("/api/register", {
      method: "POST",
      body: JSON.stringify({ email, password }),
      headers: { "Content-Type": "application/json" },
    });
    setCreatingUser(false);
    setUserCreated(true);
  }
  return (
    <section className="mt-8">
      <h1 className="text-center color text-4xl">Register</h1>
      {userCreated && (
        <div className="my-4 text-center">
          Anda sudah register. masuk ke{" "}
          <Link href={"/login"}>
            {" "}
            <span className="spanred">Login &raquo;</span>
          </Link>
        </div>
      )}
      <form className=" max-w-xs mx-auto" onSubmit={handleFormSubmit}>
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
        <button className="bg-color text-white"  type="submit" disabled={creatingUser}>
          Register
        </button>
        <div className="my-4 text-center text-gray-500">
          Or Login With
        </div>
        <button className="border border-solid border-2 flex gap-4 justify-center font-semibold">
          <Image src={"/google.png"} width={24} height={24} alt={"google"} />
          Login with google
        </button>
        <div className="text-center my-4 text-gray-500">Sudah punya akun? {' '}
        <Link className="underline" href={'/login'}>Login here</Link>
        </div>
      </form>
    </section>
  );
}
