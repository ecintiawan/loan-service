SELECT SETVAL(
	( SELECT PG_GET_SERIAL_SEQUENCE('loan', 'id') ),
	( SELECT MAX(id) FROM public.loan )
);
SELECT SETVAL(
	( SELECT PG_GET_SERIAL_SEQUENCE('investment', 'id') ),
	( SELECT MAX(id) FROM public.investment )
);
SELECT SETVAL(
	( SELECT PG_GET_SERIAL_SEQUENCE('employee', 'id') ),
	( SELECT MAX(id) FROM public.employee )
);
SELECT SETVAL(
	( SELECT PG_GET_SERIAL_SEQUENCE('borrower', 'id') ),
	( SELECT MAX(id) FROM public.borrower )
);
SELECT SETVAL(
	( SELECT PG_GET_SERIAL_SEQUENCE('investor', 'id') ),
	( SELECT MAX(id) FROM public.investor )
);