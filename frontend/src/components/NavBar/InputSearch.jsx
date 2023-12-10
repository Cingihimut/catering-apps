"use client";

import { MagnifyingGlass } from "@phosphor-icons/react";
import { useRouter } from "next/navigation";
import { useRef } from "react";

const InputSearch = () => {
  const searchRef = useRef();
  const router = useRouter();

  const handleSearch = () => {
    const keyword = searchRef.current.value.trim();
    if(keyword) {
      router.push(`/search/${keyword}`);
    }
  };

  const handleKeyDown = (event) => {
    if (event.key === "Enter") {
      event.preventDefault();
      handleSearch(btn);
    }
  };

  return (
    <div className="relative">
      <input
        id="btn"
        name="btn"
        placeholder=" cari disini"
        className="w-full p-2 rounded border border-sky-500"
        ref={searchRef}
        onKeyDown={handleKeyDown}
      />
        <button className="absolute top-2 end-2" onClick={handleSearch}>
            <MagnifyingGlass size={24} />
        </button>
    </div>
  );
};

export default InputSearch;