import Link from "next/link";
import InputSearch from "./InputSearch";

const NavBar = () => {
  return (
    <header>
      <div className="flex flex-wrap justify-between items-center mx-auto max-w-screen-xl p-4">
        <Link href="/" passHref>
          <button className="flex items-center space-x-3 rtl:space-x-reverse">
            <span className="self-center text-2xl font-semibold whitespace-nowrap">
              My Kitchen
            </span>
          </button>
        </Link>
        <div className="flex items-center space-x-6 rtl:space-x-reverse">
          <InputSearch />
          <Link href="/login" passHref>
            <button className="text-sm text-color-dark font-bold dark:text-blue-500 hover:underline">
              Login
            </button>
          </Link>
        </div>
      </div>
      <div className="flex justify-normal px-4 gap-4 text-color-third">
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
