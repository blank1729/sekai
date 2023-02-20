import React from "react";
import { Link } from "react-router-dom";

function Header() {
  return (
    <header className="absolute h-screen">
      <nav className="absolute left-0 h-[calc(100%-0.5rem)] top-0 m-1 bg-gray-800 rounded-md w-20">
        <div className="links flex flex-col gap-2 pt-28 pb-2 h-full">
          
          {/* HOME */}
          <Link
            to="/"
            className="flex flex-col justify-center bg-slate-100 mx-2 rounded-md p-1 hover:scale-110 duration-150 "
          >
            <div className="flex justify-center items-center">
              <img src="/shop-colored.svg" className="w-10 h-10 mt-1" alt="" />
            </div>
            <p className="text-center text-xs font-light mt-[2px]">HOME</p>
          </Link>

          {/* Support */}
          <Link
            to="/contact"
            className="flex flex-col justify-center bg-slate-100 mx-2 rounded-md p-1 hover:scale-110 duration-150 "
          >
            <div className="flex justify-center items-center">
              <img src="/chat-colored.svg" className="w-10 h-10 mt-1" alt="" />
            </div>
            <p className="text-center text-xs font-light mt-[2px]">SUPPORT</p>
          </Link>

          {/* Account */}
          <Link
            to="/login"
            className="mt-auto flex flex-col justify-center bg-slate-100 mx-2 rounded-md p-1 hover:scale-110 duration-150 "
          >
            <div className="flex justify-center items-center">
              <img
                src="/person-colored.svg"
                className="w-10 h-10 mt-1"
                alt=""
              />
            </div>
            <p className="text-center text-xs font-light mt-[2px]">PROFILE</p>
          </Link>
        </div>
      </nav>
    </header>
  );
}

export default Header;
