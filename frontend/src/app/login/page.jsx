// LOGIN
"use client";
import Image from "next/image";
import { useState } from "react";
import Link from "next/link";

export default function RegisterPage() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [loginInProgress, setLoginInprogress] = useState(false);

  async function handleFormSubmit(ev) {
    ev.preventDefault();
    setLoginInprogress(true);

    setLoginInprogress;
  }
  return (
    <section className="mt-8">
      <h1 className="text-center color text-4xl">Login</h1>
      <form className=" max-w-xs mx-auto" onSubmit={handleFormSubmit}>
        <input
          name="email"
          type="email"
          placeholder="email"
          value={email}
          disabled={loginInProgress}
          onChange={(ev) => setEmail(ev.target.value)}
        />
        <input
          name="password"
          type="password"
          placeholder="password"
          value={password}
          disabled={loginInProgress}
          onChange={(ev) => setPassword(ev.target.value)}
        />
        <button className="bg-color text-white" type="submit" disabled={loginInProgress}>
          Login
        </button>
        <div className="my-4 text-center text-gray-500">Or login With</div>
        <button className="border border-solid border-2 flex gap-4 justify-center font-semibold">
          <Image src={"/google.png"} width={24} height={24} alt={"google"} />
          Login with google
        </button>
      </form>
    </section>
  );
}
