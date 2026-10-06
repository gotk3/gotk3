// Same copyright and license as the rest of the files in this project

extern void goListBoxForEachFuncs(GtkListBox *box, GtkListBoxRow *row,
                                  guintptr user_data);

static inline void _gotk3_goListBoxForEachFuncs(GtkListBox *box,
                                                GtkListBoxRow *row,
                                                gpointer user_data) {
  goListBoxForEachFuncs(box, row, (guintptr)user_data);
}

static inline void _gtk_list_box_selected_foreach(GtkListBox *box,
                                                  guintptr user_data) {
  gtk_list_box_selected_foreach(box, _gotk3_goListBoxForEachFuncs,
                                (gpointer)user_data);
}
