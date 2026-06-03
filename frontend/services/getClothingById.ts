"use server";

import { cookies } from "next/headers";
import { createClient } from "@/utils/supabase/server";

export async function getClothingById(
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
        headers: {
          Authorization:
            `Bearer ${token}`,
        },
        cache: "no-store",
      }
    );

  if (!response.ok) {
  const text = await response.text();

  throw new Error(
    `Status ${response.status}: ${text}`
  );
}

  return await response.json();
}