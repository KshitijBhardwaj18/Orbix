"use client";
import React, { useState } from "react";
import { api } from "@/lib/axios";
import { registerUser } from "@/services/auth";
import toast from "react-hot-toast";
import { useRouter } from "next/navigation";
import { AxiosError } from "axios";

function SignUpCard() {
  const [showPassword, setShowPassword] = useState(false);
  const [showConfirmPassword, setShowConfirmPassword] = useState(false);
  const [password, setPassword] = useState("");
  const [email, setEmail] = useState("");
  const [username, setUsername] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [agreeToTerms, setAgreeToTerms] = useState(false);
  const router = useRouter();

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

  const onSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    try {
      const response = await registerUser({
        email: email,
        username: username,
        password: password,
      });
      console.log(response);
      toast.success("Account created successfully👍");
      router.push("/signin");
    } catch (err) {
      if (err instanceof AxiosError) {
        if (err.response) {
          // Backend responded with error (4xx or 5xx)
          if (err.response.status === 400) {
            toast.error("Invalid input. Please check your details.");
          } else if (err.response.status === 409) {
            // You can make backend return 409 Conflict if user exists
            toast.error("User already exists. Try logging in.");
          } else if (err.response.status === 500) {
            toast.error("Server error. Please try again later.");
          } else {
            toast.error(err.response.data.error || "Something went wrong.");
          }
        } else if (err.request) {
          // Request made but no response
          toast.error("No response from server. Check your connection.");
        } else {
          // Something else went wrong
          toast.error("Unexpected error occurred.");
        }
      } else {
        toast.error("Unexpected error occured");
      }
    }
  };

  return (
    <div className="mx-4 w-full max-w-md">
      <div className="bg-primary rounded-2xl border-[0.5px] border-neutral-800 p-8 shadow-2xl">
        {/* Header */}

        <div className="mr-8 flex items-center justify-center">
          <img alt="logo" src="./logo.png" className="mt-1 size-20"></img>
          <p className="text-center text-3xl font-bold text-white">Orbix</p>
        </div>

        <div className="flex items-center justify-center">
          <p className="text-written text-2xl font-[500]">Create Account</p>
        </div>

        {/* Form */}
        <form className="space-y-6" onSubmit={onSubmit}>
          <div>
            <label className="mb-2 block text-sm font-medium text-gray-300">
              Username
            </label>
            <input
              type="text"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              className="bg-secondary w-full rounded-xl border border-neutral-800 px-4 py-3 text-gray-300 placeholder-gray-400 transition-colors focus:border-blue-500 focus:outline-none"
              placeholder="Enter your username"
            />
          </div>
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

            {/* Password Strength Indicator */}
            <div className="mx-5 mt-2 flex space-x-1">
              {[...Array(5)].map((_, i) => (
                <div
                  key={i}
                  className={`h-1 flex-1 rounded ${
                    i < passwordStrength ? "bg-green-500" : "bg-gray-600"
                  }`}
                ></div>
              ))}
            </div>
          </div>

          {/* Confirm Password Field */}
          <div>
            <label className="mb-2 block text-sm font-medium text-gray-300">
              Confirm Password
            </label>
            <div className="relative">
              <input
                type={showConfirmPassword ? "text" : "password"}
                value={confirmPassword}
                onChange={(e) => setConfirmPassword(e.target.value)}
                className="bg-secondary w-full rounded-xl border border-neutral-800 px-4 py-3 pr-12 text-gray-300 placeholder-gray-400 transition-colors focus:border-blue-500 focus:outline-none"
                placeholder="Confirm your password"
              />
              <button
                type="button"
                onClick={() => setShowConfirmPassword(!showConfirmPassword)}
                className="absolute top-1/2 right-3 -translate-y-1/2 transform text-gray-400 hover:text-gray-300"
              >
                {showConfirmPassword ? (
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

          {/* Terms and Conditions */}
          <div className="flex items-start space-x-3">
            <input
              type="checkbox"
              id="terms"
              checked={agreeToTerms}
              onChange={(e) => setAgreeToTerms(e.target.checked)}
              className="mt-1 h-4 w-4 rounded border-gray-600 bg-gray-700 text-blue-500"
            />
            <label htmlFor="terms" className="text-sm text-gray-300">
              By signing up, I agree to the{" "}
              <a
                href="#"
                className="text-blue-500 underline hover:text-blue-400"
              >
                User Agreement
              </a>{" "}
              and{" "}
              <a
                href="#"
                className="text-blue-500 underline hover:text-blue-400"
              >
                Privacy Policy
              </a>
              .
            </label>
          </div>

          {/* Sign Up Button */}
          <button
            type="submit"
            className="w-full rounded-xl bg-gray-300 px-4 py-3 font-semibold text-gray-800 transition-colors hover:bg-gray-200"
          >
            Sign up
          </button>
        </form>
      </div>
    </div>
  );
}

export default SignUpCard;
