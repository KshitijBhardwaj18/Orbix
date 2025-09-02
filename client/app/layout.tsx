import type { Metadata } from "next";
import {Inter} from "next/font/google"
import "./globals.css";
import Navbar from "@/components/Navbar";

const inter = Inter({
  subsets: ['latin'],
  display: 'swap'
})

export const metadata: Metadata = {
  title: "Orbix",
  description: "Trade crypto seemlessly",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body
        className={inter.className}
      >
        <Navbar/>
        {children}
      </body>
    </html>
  );
}
