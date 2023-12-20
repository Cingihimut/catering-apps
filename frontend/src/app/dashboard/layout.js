import "../globals.css";
import DashboardNavbar from "../../components/Navbar/DashboardNavbar";


export const metadata = {
  title: "Dashboard",
  description: "Owner Dashboard",
};

export default function DashboardLayout({ children }) {
  return (
      <body suppressHydrationWarning >
        <DashboardNavbar />

        <div className="container">
          <div className="wrapper">{children}</div>
        </div>
      </body>
  );
}
