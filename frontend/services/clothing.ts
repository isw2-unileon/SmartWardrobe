"use server";

import { cookies } from "next/headers";
import { createClient } from "@/utils/supabase/server";

export async function saveClothingItem({
  imageUrl,
  colorId,
  styleId,
  typeId,
}: {
  imageUrl: string;
  colorId: number;
  styleId: number;
  typeId: number;
}) {
  const cookieStore = await cookies();
  const supabase = createClient(cookieStore);

  const {
    data: { user },
  } = await supabase.auth.getUser();

  if (!user) {
    throw new Error("User not authenticated");
  }

  const { error } = await supabase.from("clothing_items").insert({
    image_url: imageUrl,
    color_id: colorId,
    style_id: styleId,
    type_id: typeId,
    user_id: user.id,
  });

  if (error) {
    throw new Error(error.message);
  }
}
