"use client";

import { useActionState } from "react";
import { login } from "@/services/auth";

export default function LoginForm({ errorMessage }: { errorMessage?: string }) {
  const [state, formAction, isPending] = useActionState(login, { error: "" });

  return (
    // The action automatically sends all inputs that have the "name" tag
    <form action={formAction}>
      <h2>Login</h2>

      {/* The error is displayed if the prop receives it */}
      {state?.error && (
        <div style={{ padding: '10px', backgroundColor: '#ffebee', color: '#c62828', borderRadius: '4px' }}>
          {state.error}
        </div>
      )}

      <input
        type="email"
        name="email"    
        placeholder="Email"
        required
      />

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
  );
}