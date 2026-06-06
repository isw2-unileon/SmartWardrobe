"use server";

import { cookies } from "next/headers";
import { createClient } from "@/utils/supabase/server";

export interface ClipPredictionResponse {
  color: string;
  style: string;
  type: string;
}

export async function analyzeClothing(
  file: File,
): Promise<ClipPredictionResponse> {
  const cookieStore = await cookies();

  const supabase = createClient(cookieStore);

  const {
    data: { session },
  } = await supabase.auth.getSession();

  const token = session?.access_token;

  const formData = new FormData();

  formData.append("image", file);

  const response = await fetch(
    `${process.env.NEXT_PUBLIC_API_URL}/api/clothing/analyze`,
    {
      method: "POST",

      headers: {
        Authorization: `Bearer ${token}`,
      },

      body: formData,
    },
  );

  if (!response.ok) {
    const errorText = await response.text();

    console.log("CLIP ERROR:", errorText);

    throw new Error(errorText);
  }

  return await response.json();
}
