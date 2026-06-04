"use server";

import { cookies } from "next/headers";
import { createClient } from "@/utils/supabase/server";

type UpdateParams = {
  id: number;

  typeId: number;
  typeName: string;

  colorId: number;
  colorName: string;

  styleId: number;
  styleName: string;

  imageUrl: string;
};

export async function updateClothing({
  id,
  typeId,
  typeName,
  colorId,
  colorName,
  styleId,
  styleName,
  imageUrl,
}: UpdateParams) {
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
    `${process.env.NEXT_PUBLIC_API_URL}/api/clothingItem/${id}`,
    {
      method: "PUT",

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

    throw new Error(error.error || "Failed to update clothing");
  }

  return await response.json();
}
