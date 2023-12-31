"use client";

import { useEffect, useState } from "react";
import Link from "next/link";
import Icons from "./component/icons/page";
import Menu from "./component/menu";
import Image from "next/image";
import Navbar from "./component/navbar";
import { useUserStore } from "@/stores/userStore";

const Home = () => {
  const [showScrollButton, setShowScrollButton] = useState(false);
  const { user } = useUserStore();
  console.log(user);
  useEffect(() => {
    if (typeof window !== "undefined") {
      const handleScroll = () => {
        const shouldShowButton = window.scrollY > 100;
        setShowScrollButton(shouldShowButton);
      };

      window.addEventListener("scroll", handleScroll);

      return () => {
        window.removeEventListener("scroll", handleScroll);
      };
    }
  }, []);

  useEffect(() => {
    const snapScprit = "https://app.sandbox.midtrans.com/snap/snap.js";
    const clientKey = process.env.NEXT_PUBLIC_CLIENT;
    const script = document.createElement("script");
    script.src = snapScprit;
    script.setAttribute("data-client-key", clientKey);
    script.async = true;

    document.body.appendChild(script);

    return () => {
      document.body.removeChild(script);
    };
  }, []);

  const scrollToTop = () => {
    window.scrollTo({ top: 0, behavior: "smooth" });
  };

  return (
    <section>
      <Navbar />
      <div className="flex justify-between items-center mt-8">
        <div className="w-1/2 pr-8">
          <h1 className="text-4xl font-bold mb-4">
            Mama arsy <span className="color">Catering</span> is here
          </h1>
          <p className="text-lg">
            It is a long established fact that a reader will be distracted by
            the readable content of a page when looking at its layout. The point
            of using Lorem Ipsum is that it has a more-or-less
          </p>
          <div className="flex gap-2 text-sm mt-8">
            <Link href={"#halamanMenu"}>
              <button className="bg-color flex gap-2 text-white px-9 py-auto items-center rounded-full justify-center uppercase">
                Order Now
                <Icons />
              </button>
            </Link>
            <button className="flex gap-2 py-2 text-gray-500 font-semibold">
              Learn More
              <Icons />
            </button>
          </div>
        </div>

        <div className="w-1/2">
          <Image
            src="/logo.png"
            alt="Makanan Produk"
            className="w-full h-auto rounded-full"
            width={400}
            height={200}
          />
        </div>
      </div>
      <div id="halamanMenu">
        <Menu />
      </div>

      <div className="text-center font-semibold color text-4xl mt-16">
        <h1 className="">About me</h1>
      </div>
      <p className="text-center mb-4 mx-16 text-gray-500">
        Lorem Ipsum is simply dummy text of the printing and typesetting
        industry. Lorem Ipsum has been the industry's standard dummy text ever
        since the 1500s, when an unknown printer took a galley of type and
        scrambled it to make a type specimen book.
      </p>
      <div className="text-center font-semibold color text-4xl mt-16">
        <h1 className="">Contact</h1>
        <div id="contactMe">
          <Link
            href={"/"}
            className="text-center mb-4 mx-16 text-2xl text-gray-400"
          >
            +62000000x
          </Link>
        </div>
      </div>

      {showScrollButton && (
        <button
          className="fixed bottom-2 right-8 bg-color text-white rounded-full p-2 w-10 h-10 flex items-center justify-center"
          onClick={scrollToTop}
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            strokeWidth={1.5}
            stroke="currentColor"
            dataSlot="icon"
            className="w-6 h-6"
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              d="M8.25 6.75 12 3m0 0 3.75 3.75M12 3v18"
            />
          </svg>
        </button>
      )}
    </section>
  );
};

export default Home;