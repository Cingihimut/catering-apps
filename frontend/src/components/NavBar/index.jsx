import Link from "next/link";
import InputSearch from "./InputSearch";

const NavBar = () => {
  return (
    <header>
      <div className="flex flex-wrap justify-between items-center mx-auto max-w-screen-xl p-4 border-b-2">
        <Link href="/" passHref className="text-color-blackMode">
          <button className="flex items-center space-x-3 rtl:space-x-reverse">
            <span className="self-center text-2xl font-semibold whitespace-nowrap">
              My Kitchen
            </span>
          </button>
        </Link>
        <div className="flex items-center space-x-6 rtl:space-x-reverse">
          <InputSearch />
          <Link href="/login" passHref className="border border-2 border-color-third rounded hover:bg-color-third">
            <button className="text-sm text-color-third hover:text-color-white font-bold dark:text-blue-500 p-1">
              Masuk
            </button>
          </Link>
          <Link href="/register" passHref>
            <button className="text-sm text-color-third font-bold dark:text-blue-500">
              Daftar
            </button>
          </Link>
        </div>
      </div>
      <div className="flex justify-normal px-4 gap-4 text-color-monocrome text-sm flex justify-center p-2">
        <Link href="../Menu/rekomendasi" passHref>
          <button className="hover:text-color-dark transition duration-300 ease-in-out">
            Rekomendasi kami
          </button>
        </Link>
        <Link href="../Menu/terbaik" passHref>
          <button className="hover:text-color-dark transition duration-300 ease-in-out">
            Pilihan Terbaik
          </button>
        </Link>
        <Link href="../Menu/terlaris" passHref>
          <button className="hover:text-color-dark transition duration-300 ease-in-out">
            Pilihan Terlaris
          </button>
        </Link>
        <Link href="/Menu" passHref>
          <button className="hover:text-color-dark transition duration-300 ease-in-out">
            Menu lengkap
          </button>
        </Link>
      </div>
    </header>
  );
};

export default NavBar;
