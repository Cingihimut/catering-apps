"use client";
import React from "react";
import SwaggerUIReact from "swagger-ui-react";
import "swagger-ui-react/swagger-ui.css";

const SwaggerUI = () => {
  return (
    <div>
      <SwaggerUIReact url="/api/api-spec.json" />;
    </div>
  );
};

export default SwaggerUI;
