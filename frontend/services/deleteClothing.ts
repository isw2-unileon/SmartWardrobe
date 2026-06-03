"use server";

import { cookies } from "next/headers";
import { createClient } from "@/utils/supabase/server";

export async function deleteClothing(
  id: number
) {
  const cookieStore =
    await cookies();

  const supabase =
    createClient(cookieStore);

  const {
    data: { session },
  } =
    await supabase.auth.getSession();

  const token =
    session?.access_token;

  if (!token) {
    throw new Error(
      "No session found"
    );
  }

  const response =
    await fetch(
      `${process.env.NEXT_PUBLIC_API_URL}/api/clothingItem/${id}`,
      {
        method: "DELETE",

        headers: {
          Authorization:
            `Bearer ${token}`,
        },
      }
    );

  if (!response.ok) {
    const error =
      await response.json();

    throw new Error(
      error.error ||
        "Failed to delete clothing"
    );
  }

  return await response.json();
}