"use client";

import { useActionState } from "react";
import { login } from "@/services/auth";

export default function LoginForm() {
  const [state, formAction, isPending] = useActionState(login, { error: "" });

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
        <form action={formAction}>
          <h2
            style={{
              textAlign: "center",
              marginBottom: "1rem",
            }}
          >
            Login
          </h2>

          {state?.error && <div className="error-message">{state.error}</div>}

          <input type="email" name="email" placeholder="Email" required />

          <input
            type="password"
            name="password"
            placeholder="Password"
            required
          />

          <button type="submit" disabled={isPending}>
            {isPending ? "Loading..." : "Login"}
          </button>
        </form>
      </div>
    </div>
  );
}
