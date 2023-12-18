import React from "react";
import DashboardTabs from "../../components/Dashboard/Tabs";
const Dashboard = () => {
  return (
    <main className="mt-4 h-full">
      <div className="flex justify-center w-full">
        <DashboardTabs />
      </div>
    </main>
  );
};

export default Dashboard;
