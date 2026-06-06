"use server";

import { cookies } from "next/headers";
import { createClient } from "@/utils/supabase/server";

type GenerateOutfitParams = {
  city: string;

  country: string;

  startDate: string;

  endDate: string;
};

export async function generateOutfit({
  city,
  country,
  startDate,
  endDate,
}: GenerateOutfitParams) {
  const cookieStore = await cookies();

  const supabase = createClient(cookieStore);

  const {
    data: { session },
  } = await supabase.auth.getSession();

  const token = session?.access_token;

  if (!token) {
    throw new Error("No session found");
  }

  const response = await fetch(
    `${process.env.NEXT_PUBLIC_API_URL}/api/generateOutfit/days`,
    {
      method: "POST",

      headers: {
        "Content-Type": "application/json",

        Authorization: `Bearer ${token}`,
      },

      body: JSON.stringify({
        city,
        country,
        start_date: startDate,

        end_date: endDate,
      }),
    },
  );

  if (!response.ok) {
    const error = await response.json();

    throw new Error(error.error || "Failed to generate outfit");
  }

  return await response.json();
}
