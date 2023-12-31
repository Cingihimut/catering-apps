import { Inter } from "next/font/google";
import "./globals.css";
import Navbar from "./component/navbar";

const inter = Inter({ subsets: ["latin"] });

export const metadata = {
  title: "Create Next App",
  description: "Generated by create next app",
};

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <div className="container">
          <div className="wrapper">
            {children}
          </div>
          <footer className="border-t p-8 text-center text-gray-500">
            &copy; 2023
          </footer>
        </div>
      </body>
    </html>
  );
}
