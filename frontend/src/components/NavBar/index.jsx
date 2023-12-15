import Link from "next/link";
import InputSearch from "./InputSearch";
import UserActionButton from "./UserActionButton";
const NavBar = () => {
  return (
    <header>
      <div className="flex flex-wrap justify-between items-center mx-auto max-w-screen-xl p-4 border-b-2">
        <Link href="/" passHref className="text-color-third">
          <button className="flex items-center space-x-3 rtl:space-x-reverse">
            <span className="self-center text-2xl font-semibold whitespace-nowrap">
              My Kitchen
            </span>
          </button>
        </Link>
        <div className="flex items-center space-x-6 rtl:space-x-reverse ml-auto">
          <InputSearch />
          <UserActionButton />
        </div>
      </div>
      <div className="flex justify-end">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          strokeWidth={1.5}
          stroke="currentColor"
          className="w-6 h-6"
        >
          <path
            strokeLinecap="round"
            strokeLinejoin="round"
            d="M15 10.5a3 3 0 11-6 0 3 3 0 016 0z"
          />
          <path
            strokeLinecap="round"
            strokeLinejoin="round"
            d="M19.5 10.5c0 7.142-7.5 11.25-7.5 11.25S4.5 17.642 4.5 10.5a7.5 7.5 0 1115 0z"
          />
        </svg>
      </div>
    </header>
  );
};

export default NavBar;
