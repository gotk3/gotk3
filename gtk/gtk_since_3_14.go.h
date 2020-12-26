// Same copyright and license as the rest of the files in this project

extern void goListBoxForEachFuncs(GtkListBox *box, GtkListBoxRow *row,
                                  gpointer user_data);

static inline void _gtk_list_box_selected_foreach(GtkListBox *box,
                                                  gpointer user_data) {
  gtk_list_box_selected_foreach(
      box, (GtkListBoxForeachFunc)(goListBoxForEachFuncs), user_data);
}
