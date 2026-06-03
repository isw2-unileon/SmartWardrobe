"use server";

import { cookies } from "next/headers";
import { createClient } from "@/utils/supabase/server";

type SaveParams = {
  typeId: number;
  colorId: number;
  styleId: number;
  imageUrl: string;

  typeName: string;
  colorName: string;
  styleName: string;
};

export async function saveClothingItem({
  typeId,
  colorId,
  styleId,
  imageUrl,
  typeName,
  colorName,
  styleName,
}: SaveParams) {
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
    `${process.env.NEXT_PUBLIC_API_URL}/api/clothingItem`,
    {
      method: "POST",

      headers: {
        "Content-Type": "application/json",

        Authorization: `Bearer ${token}`,
      },

      body: JSON.stringify({
        type: {
          id: typeId,
          name: typeName,
        },

        color: {
          id: colorId,
          name: colorName,
        },

        style: {
          id: styleId,
          name: styleName,
        },

        imageUrl,
      }),
    },
  );

  if (!response.ok) {
    const error = await response.json();

    throw new Error(error.error || "Failed to save clothing");
  }

  return await response.json();
}
