"use client";

import Kategori from "./kategori/page";
import Carousel from "./bannerPromo/Caraousel";
import useStore from "@/stores/store";
import Link from "next/link";

const Page = () => {
  const user = useStore((state) => state.user);
  console.log(user);
  return (
    <>
      <main>
        <div className="p-4">
          {user.email ?(
            <h2>Selamat Datang {user.email}</h2>
          ) : (
            <p>Selamat Datang</p>
          )}
          <Carousel />
        </div>
      </main>
      <Link href="/kategori/lengkap">
        <p className="font-bold p-4 text-2xl">Kategori</p>
      </Link>
      <Kategori />
    </>
  );
};
export default Page;
