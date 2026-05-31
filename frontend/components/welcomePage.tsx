"use client";

import { useRouter } from "next/navigation";

export default function WelcomePage() {
  const router = useRouter();

  return (
    <div
      style={{
        minHeight: "100vh",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        padding: "2rem",
      }}
    >
      <div
        style={{
          width: "100%",
          maxWidth: "420px",

          backgroundColor: "#FFFDFB",

          borderRadius: "28px",

          border: "1px solid #E5D8CC",

          padding: "2.5rem",

          boxShadow:
            "0 10px 25px rgba(0,0,0,0.05), 0 4px 10px rgba(0,0,0,0.03)",

          display: "flex",
          flexDirection: "column",
          gap: "1.5rem",
          alignItems: "center",
        }}
      >
        <h1
          style={{
            marginBottom: "1rem",
            textAlign: "center",
          }}
        >
          Smart Wardrobe
        </h1>

        <button
          style={{ width: "220px" }}
          onClick={() => router.push("/login")}
        >
          Log In
        </button>

        <button
          style={{ width: "220px" }}
          onClick={() => router.push("/register")}
        >
          Sign Up
        </button>
      </div>
    </div>
  );
}
