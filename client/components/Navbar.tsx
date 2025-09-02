import React from "react";

function Navbar() {
  const links = [
    {
      name: "Spot",
      href: "#",
    },
    {
      name: "Futures",
      href: "#",
    },
    {
      name: "Lend",
      href: "#",
    },
    {
      name: "blogs",
      href: "#",
    },
  ];
  return (
    <div className="bg-secondary w-full">
      <div className="flex flex-row items-center justify-between  p-2 px-4">
        <div className="flex flex-row gap-10">
          <div className="flex items-center justify-center">
            <img alt="logo" src="./logo.png" className="mt-1 size-8"></img>
            <p className="text-center font-bold text-white">Orbix</p>
          </div>
          <div className="flex flex-row items-center justify-between gap-8">
            {links.map((link, idx) => (
              <a key={idx} className="text-sm font-[600] text-neutral-400">
                {link.name}
              </a>
            ))}
          </div>
        </div>

        <div className="flex items-center justify-center mr-35">
          <div className="relative w-[25rem]">
            <input
              type="text"
              placeholder="Search markets"
              className="w-full rounded-xl border-none bg-neutral-800 px-2 py-1 pr-12 pl-12 text-gray-300 placeholder-gray-400 placeholder:text-sm focus:border-blue-500/50 focus:outline-none text-sm"
            />

            {/* Search Icon */}
            <div className="absolute top-1/2 left-4 -translate-y-1/2 transform">
              <svg
                className="h-5 w-5 text-gray-400"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth={2}
                  d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
                />
              </svg>
            </div>

            {/* Forward Slash Button */}
            <button className="absolute top-1/2 right-3 flex h-5 w-6 -translate-y-1/2 transform items-center justify-center rounded border border-gray-400 text-sm text-gray-400 hover:bg-gray-700/50 ">
              /
            </button>
          </div>
        </div>
        <div className="flex flex-row gap-5">
          <button className="flex items-center justify-center rounded-lg bg-green-800 p-[0.3rem] px-2 text-sm font-[600] text-green-400 cursor-pointer hover:bg-green-900" >
            Sign Up
          </button>
          <button className="flex items-center justify-center rounded-lg bg-sky-900 p-[0.3rem] px-2 text-sm font-[600] text-sky-400 cursor-pointer hover:bg-sky-950">
            Sign In
          </button>
        </div>
      </div>
    </div>
  );
}

export default Navbar;
