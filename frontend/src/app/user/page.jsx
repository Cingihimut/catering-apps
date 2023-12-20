// components/UserDashboard.js
"use client";
import NavbarUser from "@/components/Navbar/NavbarUser";
import React, { useState } from "react";

const UserDashboard = () => {
  const [formData, setFormData] = useState({
    fullName: "",
    address: "",
    phoneNumber: "",
    email: "",
  });

  const [isEditMode, setIsEditMode] = useState(true);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prevData) => ({
      ...prevData,
      [name]: value,
    }));
  };

  const handleSave = () => {
    // Handle save logic here (e.g., send data to server)
    console.log("Saving data:", formData);
    // You can implement your API call or save logic here

    // Toggle back to edit mode
    setIsEditMode(true);
  };

  const handleEdit = () => {
    // Toggle to edit mode
    setIsEditMode(false);
  };

  return (
    <div className="sticky top-0 z-10" style={{ height: "70px" }}>
     <NavbarUser/>
      <div className="flex justify-between mt-8">
        {/* Left Card - User Info */}
        <div className="w-1/2 p-4">
          <div className="bg-white p-6 rounded shadow-md">
            <h2 className="text-xl font-semibold mb-4">User Information</h2>
            {isEditMode ? (
              <form>
                <div className="mb-4">
                  <label htmlFor="fullName" className="block text-sm font-medium text-gray-600">
                    Full Name:
                  </label>
                  <input
                    type="text"
                    id="fullName"
                    name="fullName"
                    className="border rounded px-3 py-2 w-full"
                    placeholder="John Doe"
                    value={formData.fullName}
                    onChange={handleChange}
                  />
                </div>
                <div className="mb-4">
                  <label htmlFor="address" className="block text-sm font-medium text-gray-600">
                    Address:
                  </label>
                  <input
                    type="text"
                    id="address"
                    name="address"
                    className="border rounded px-3 py-2 w-full"
                    placeholder="123 Main Street, Cityville"
                    value={formData.address}
                    onChange={handleChange}
                  />
                </div>
                <div className="mb-4">
                  <label htmlFor="phoneNumber" className="block text-sm font-medium text-gray-600">
                    Phone Number:
                  </label>
                  <input
                    type="tel"
                    id="phoneNumber"
                    name="phoneNumber"
                    className="border rounded px-3 py-2 w-full"
                    placeholder="123-456-7890"
                    value={formData.phoneNumber}
                    onChange={handleChange}
                  />
                </div>
                <div className="mb-4">
                  <label htmlFor="email" className="block text-sm font-medium text-gray-600">
                    Email:
                  </label>
                  <input
                    type="email"
                    id="email"
                    name="email"
                    className="border rounded px-3 py-2 w-full"
                    placeholder="john.doe@example.com"
                    value={formData.email}
                    onChange={handleChange}
                  />
                </div>
                <button
                  type="button"
                  onClick={handleSave}
                  className="bg-blue-500 text-white px-4 py-2 rounded-full hover:bg-blue-600"
                >
                  Save Profile
                </button>
              </form>
            ) : (
              <div>
                <div className="mb-4">
                  <label htmlFor="fullName" className="block text-sm font-medium text-gray-600">
                    Full Name:
                  </label>
                  <p className="text-gray-800">{formData.fullName}</p>
                </div>
                <div className="mb-4">
                  <label htmlFor="address" className="block text-sm font-medium text-gray-600">
                    Address:
                  </label>
                  <p className="text-gray-800">{formData.address}</p>
                </div>
                <div className="mb-4">
                  <label htmlFor="phoneNumber" className="block text-sm font-medium text-gray-600">
                    Phone Number:
                  </label>
                  <p className="text-gray-800">{formData.phoneNumber}</p>
                </div>
                <div className="mb-4">
                  <label htmlFor="email" className="block text-sm font-medium text-gray-600">
                    Email:
                  </label>
                  <p className="text-gray-800">{formData.email}</p>
                </div>
                <button
                  type="button"
                  onClick={handleEdit}
                  className="bg-yellow-500 text-white px-4 py-2 rounded-full hover:bg-yellow-600"
                >
                  Edit Profile
                </button>
              </div>
            )}
          </div>
        </div>

        <div className="w-1/2 p-4">
          <div className="bg-white p-6 rounded shadow-md">
            <h2 className="text-xl font-semibold mb-4">Transaction History</h2>
            <div className="mb-4">
              <p className="text-gray-800">Transaction 1: Success</p>
            </div>
            <div className="mb-4">
              <p className="text-gray-800">Transaction 2: Pending</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default UserDashboard;
