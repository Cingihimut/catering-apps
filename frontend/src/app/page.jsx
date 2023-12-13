"use client";
import Kategori from "./kategori/page";
import Carousel from "./bannerPromo/Caraousel";
import useStore from "@/stores/store";

const Page = () => {
  const user = useStore((state) => state.user);
  console.log(user);
  return (
    <>
      <main>
        <div className="py-6">
          {user.email ? (
            <h2>Selamat Datang {user.email}</h2>
          ) : (
            <p>Selamat Datang</p>
          )}
          <Carousel />
        </div>
      </main>
      <p className="font-bold p-4 text-2xl">Kategori</p>
      <Kategori />
    </>
  );
};
export default Page;
