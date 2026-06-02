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
    data: { user },
  } = await supabase.auth.getUser();

  if (!user) {
    throw new Error(
      "Not authenticated"
    );
  }

  const { error } =
    await supabase
      .from("clothing_items")
      .delete()
      .eq("id", id)
      .eq("user_id", user.id);

  if (error) {
    throw new Error(
      error.message
    );
  }
}