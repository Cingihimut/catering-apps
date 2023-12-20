// LOGIN
"use client";
import Image from "next/image";
import { useState } from "react";
import { useRouter } from "next/navigation";
import { useUserStore } from "@/stores/userStore";
import decodeToken from "../../utils/decodeToken";

export default function RegisterPage() {
  const { user, setUser } = useUserStore();
  const router = useRouter();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [loginInProgress, setLoginInprogress] = useState(false);
  const API_URL = process.env.API_URL + "/api/users/register";

  async function handleFormSubmit(ev) {
    ev.preventDefault();
    const response = await fetch(API_URL, {
      method: "POST",
      body: JSON.stringify({ email, password }),
      headers: {
        "Content-Type": "application/json",
      },
    });
    if (response.ok) {
      const data = await response.json();
      const decodedUser = decodeToken(data.token);
      window.localStorage.setItem("token", data.token);
      setUser(decodedUser);
      if (decodedUser.role == "User") {
        router.push("/");
      } else {
        router.push("/dashboard");
      }
    }
    setLoginInprogress(true);
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
        <button
          className="bg-color text-white"
          type="submit"
          disabled={loginInProgress}
        >
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