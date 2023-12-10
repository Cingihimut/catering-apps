import "@/app/globals.css";
import { Inter  } from "next/font/google";

const inter = Inter ({ 
  subsets: ["latin"],
  display: 'swap'
});

export const metadata = {
  title: "Catering apps",
  description: "Pesan makanan anda dan nikmati bersama orang tersayang",
};

export default function RootLayout({ children }) {
  return (
    <html lang="en" className={inter.className}>
      <body>{children}</body>
    </html>
  );
}