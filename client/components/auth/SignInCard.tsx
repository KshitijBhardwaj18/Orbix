"use client";
import React, { useState } from "react";

function SignInCard() {
  const [showPassword, setShowPassword] = useState(false);
  const [showConfirmPassword, setShowConfirmPassword] = useState(false);
  const [password, setPassword] = useState("");
  const [email, setEmail] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [agreeToTerms, setAgreeToTerms] = useState(false);

  const passwordStrengthLabel = ["poor", "bad", "Okayish", "Good"];

  // Password strength indicator
  const getPasswordStrength = (password: string) => {
    if (password.length === 0) return 0;
    if (password.length < 4) return 1;
    if (password.length < 6) return 2;
    if (password.length < 8) return 3;
    if (password.length < 10) return 4;
    return 5;
  };

  const passwordStrength = getPasswordStrength(password);

  return (
    <div className="mx-4 w-full max-w-md">
      <div className="bg-primary rounded-2xl border-[0.5px] border-neutral-800 p-8 shadow-2xl">
        {/* Header */}

        <div className="mr-8 flex items-center justify-center ml-4">
          <img alt="logo" src="./logo.png" className="mt-1 size-20"></img>
    
        </div>

        <div className="flex items-center justify-center">
          <p className="text-written text-2xl font-[500]">Create Account</p>
        </div>

        {/* Form */}
        <form className="space-y-6">
          {/* Email Field */}
          <div>
            <label className="mb-2 block text-sm font-medium text-gray-300">
              Email
            </label>
            <input
              type="email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              className="bg-secondary w-full rounded-xl border border-neutral-800 px-4 py-3 text-gray-300 placeholder-gray-400 transition-colors focus:border-blue-500 focus:outline-none"
              placeholder="Enter your email"
            />
          </div>

          {/* Password Field */}
          <div>
            <label className="mb-2 block text-sm font-medium text-gray-300">
              Password
            </label>
            <div className="relative">
              <input
                type={showPassword ? "text" : "password"}
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                className="bg-secondary w-full rounded-xl border border-neutral-800 px-4 py-3 pr-12 text-gray-300 placeholder-gray-400 transition-colors focus:border-blue-500 focus:outline-none"
                placeholder="Enter your password"
              />
              <button
                type="button"
                onClick={() => setShowPassword(!showPassword)}
                className="absolute top-1/2 right-3 -translate-y-1/2 transform text-gray-400 hover:text-gray-300"
              >
                {showPassword ? (
                  <svg
                    className="h-5 w-5"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth={2}
                      d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.878 9.878L3 3m6.878 6.878L21 21"
                    />
                  </svg>
                ) : (
                  <svg
                    className="h-5 w-5"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth={2}
                      d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"
                    />
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth={2}
                      d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"
                    />
                  </svg>
                )}
              </button>
            </div>
          </div>

         

          <button
            type="submit"
            className="w-full rounded-xl bg-gray-300 px-4 py-3 font-semibold text-gray-800 transition-colors hover:bg-gray-200 cursor-pointer"
          >
            Sign In
          </button>
        </form>
      </div>
    </div>
  );
}

export default SignInCard;
