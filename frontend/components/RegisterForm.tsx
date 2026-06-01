"use client";

import { useActionState, useState } from "react";
import { signUp } from "@/services/auth";

export default function RegisterForm() {
  const [state, formAction, isPending] = useActionState(signUp, {
    error: "",
  });

  const [passwordError, setPasswordError] = useState("");

  function validatePasswords(formData: FormData) {
    const password = formData.get("password") as string;
    const confirmPassword = formData.get("confirmPassword") as string;

    if (password.length < 8) {
      setPasswordError("Password must be at least 8 characters long");
      return false;
    }

    if (password !== confirmPassword) {
      setPasswordError("Passwords do not match");
      return false;
    }

    setPasswordError("");
    return true;
  }

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
          maxWidth: "450px",
          backgroundColor: "#FFFDFB",
          border: "1px solid #E5D8CC",
          borderRadius: "24px",
          padding: "2rem",
          boxShadow:
            "0 10px 25px rgba(0,0,0,0.05), 0 4px 10px rgba(0,0,0,0.03)",
        }}
      >
        <form
          action={(formData) => {
            if (validatePasswords(formData)) {
              formAction(formData);
            }
          }}
          style={{
            display: "flex",
            flexDirection: "column",
            gap: "1.2rem",
          }}
        >
          <h2
            style={{
              textAlign: "center",
              marginBottom: "1.5rem",
            }}
          >
            Sign Up
          </h2>

          {state?.error && <div className="error-message">{state.error}</div>}

          {passwordError && (
            <div className="error-message">{passwordError}</div>
          )}

          <input type="email" name="email" placeholder="Email" required />

          <input
            type="password"
            name="password"
            placeholder="Password"
            required
          />

          <input
            type="password"
            name="confirmPassword"
            placeholder="Confirm password"
            required
          />

          <button
            type="submit"
            disabled={isPending}
            style={{
              marginTop: "1rem",
            }}
          >
            {isPending ? "Creating account..." : "Sign Up"}
          </button>
        </form>
      </div>
    </div>
  );
}
