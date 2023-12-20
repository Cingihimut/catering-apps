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

  async function handleFormSubmit(ev) {
    ev.preventDefault();
    const response = await fetch(`https://labs.mhdaris.me/api/users/login`, {
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
<<<<<<< HEAD
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
=======
    <div className="min-h-screen flex items-center justify-center bg-white">
      <div className="bg-gray-100 p-9 rounded-lg shadow-lg">
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
          <p className="text-sm text-gray-600">
            Masuk lewat{" "}
            <a href="/signin" className="text-blue-500 hover:underline">
              github?{" "}
            </a>
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
>>>>>>> main
  );
}
