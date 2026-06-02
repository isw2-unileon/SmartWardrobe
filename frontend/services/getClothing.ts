"use server";

import { cookies } from "next/headers";
import { createClient } from "@/utils/supabase/server";

export async function getClothingItems() {
  const cookieStore = await cookies();
  const supabase =
    createClient(cookieStore);

  const {
    data: { user },
  } = await supabase.auth.getUser();

  if (!user) {
    return [];
  }

  const { data, error } =
    await supabase
      .from("clothing_items")
      .select("*")
      .eq("user_id", user.id)
      .order("id", {
        ascending: false,
      });

  if (error) {
    throw new Error(error.message);
  }

  return data;
}