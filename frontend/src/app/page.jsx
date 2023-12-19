"use client"
import React, { useEffect } from "react";
import Link from "next/link";
import Icons from "./component/icons/page";
import Menu from "./component/menu";
import Image from "next/image";


export default function Home() {

  useEffect(() => {
    const snapScrpit = "https://app.sandbox.midtrans.com/snap/snap.js"
    const clientKey  = process.env.NEXT_PUBLIC_CLIENT
    const script = document.createElement('script')
    script.src = snapScrpit
    script.setAttribute('data-client-key', clientKey)
    script.async = true

    document.body.appendChild(script)
    return() => {
      document.body.removeChild(script)
    }

  }, []);

  return (
    <section>
      <div className="flex justify-between items-center mt-8">
        <div className="w-1/2 pr-8">
          <h1 className="text-4xl font-bold mb-4">
            My moms <span className="color">Catering</span> is here
          </h1>
          <p className="text-lg">
            It is a long established fact that a reader will be distracted by
            the readable content of a page when looking at its layout. The point
            of using Lorem Ipsum is that it has a more-or-less
          </p>
          <div className="flex gap-2 text-sm mt-8">
            <Link href={"/"}>
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
      <Menu />
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

        <Link
          href={"/"}
          className="text-center mb-4 mx-16 text-2xl text-gray-400"
        >
          +62000000
        </Link>
      </div>
    </section>
  );
}
