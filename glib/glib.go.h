/*
 * Copyright (c) 2013 Conformal Systems <info@conformal.com>
 *
 * This file originated from: http://opensource.conformal.com/
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

#include <stdint.h>
#include <stdlib.h>
#include <stdio.h>

/* GObject Type Casting */
static GObject *
toGObject(void *p)
{
	return (G_OBJECT(p));
}

static GType
_g_type_from_instance(gpointer instance)
{
	return (G_TYPE_FROM_INSTANCE(instance));
}

/* Wrapper to avoid variable arg list */
static void
_g_object_set_one(gpointer object, const gchar *property_name, void *val)
{
	g_object_set(object, property_name, *(gpointer **)val, NULL);
}

static GValue *
alloc_gvalue_list(int n)
{
	GValue		*valv;

	valv = g_new0(GValue, n);
	return (valv);
}

static void
val_list_insert(GValue *valv, int i, GValue *val)
{
	valv[i] = *val;
}

typedef struct {
	int		func_n;
	gboolean	ret;
	guint		id;
} idleinfo;

/* Call to Go to nil this context to free memory next garbage collection */
extern void _go_nil_unused_idle_ctx(int n);

/*
 * Called after a function is removed from the main loop context to free the
 * idle context.
 */
static void
idleinfo_free(gpointer data)
{
	idleinfo	*idl = (idleinfo *)data;

	_go_nil_unused_idle_ctx(idl->func_n);
	free(idl);
}

/* Call to Go to run func in the main loop context */
extern void _go_glib_idle_fn(idleinfo *idl);

static gboolean
_g_idle_run(gpointer user_data)
{
	idleinfo	*idl = (idleinfo *)user_data;

	_go_glib_idle_fn(idl);
	return (idl->ret);
}

/*
 * Create idleinfo context and add _g_idle_run and its context to run
 * in the GTK main loop during idle state
 */
static idleinfo *
_g_idle_add(int func_n)
{
	idleinfo	*idl;

	idl = (idleinfo *)malloc(sizeof(idleinfo));
	idl->func_n = func_n;
	idl->id = g_idle_add_full(G_PRIORITY_DEFAULT_IDLE, _g_idle_run,
	    (gpointer)idl, idleinfo_free);
	return (idl);
}

/*
 * GValue
 */

static GValue *
_g_value_alloc()
{
	return (g_new0(GValue, 1));
}

static GValue *
_g_value_init(GType g_type)
{
	GValue          *value;

	value = g_new0(GValue, 1);
	return (g_value_init(value, g_type));
}

static gboolean
_g_is_value(GValue *val)
{
	return (G_IS_VALUE(val));
}

static GType
_g_value_type(GValue *val)
{
	return (G_VALUE_TYPE(val));
}

static GType
_g_value_fundamental(GType type)
{
	return (G_TYPE_FUNDAMENTAL(type));
}

/*
 * Closure support
 */

extern void	goMarshal(GClosure *, GValue *, guint, GValue *, gpointer, GValue *);//gpointer);

static GClosure *
_g_closure_new()
{
	GClosure	*closure;

	closure = g_closure_new_simple(sizeof(GClosure), NULL);
	g_closure_set_marshal(closure, (GClosureMarshal)(goMarshal));
	return closure;
}

static GClosure *
_g_closure_new_with_data(gpointer marshal_data)
{
	GClosure	*closure;

	closure = g_closure_new_simple(sizeof(GClosure), NULL);
	g_closure_set_meta_marshal(closure, marshal_data,
	    (GClosureMarshal)(goMarshal));
	return closure;
}
