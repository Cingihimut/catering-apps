import { Gabarito } from "next/font/google";
import "@/app/globals.css";
import NavBar from "@/components/NavBar/index";
import Footer from "@/components/Footer";

const gabarito = Gabarito({ subsets: ["latin"] });

export const metadata = {
  title: "Catering apps",
  description: "Pesan makanan anda dan nikmati bersama orang tersayang",
};

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body className={`${gabarito.className}`} suppressHydrationWarning={true}>
        <NavBar/>
        {children}
        <Footer/>
      </body>
    </html>
  );
}