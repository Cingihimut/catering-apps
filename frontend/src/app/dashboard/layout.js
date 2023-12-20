import { Inter } from "next/font/google";
import "../globals.css";
import DashboardNavbar from "../../components/Navbar/DashboardNavbar";

const inter = Inter({ subsets: ["latin"] });

export const metadata = {
  title: "Dashboard",
  description: "Owner Dashboard",
};

export default function DashboardLayout({ children }) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <DashboardNavbar />

        <div className="container">
          <div className="wrapper">{children}</div>
        </div>
      </body>
    </html>
  );
}
