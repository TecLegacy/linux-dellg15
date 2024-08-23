log.Printf("2 registration_repository.go :GetCompanyDomain")
	resp := &grpcapi_mussad_svc_tenant.TradeActivityResponse{}

	tradeActivityInfoMap := make(map[int32]*grpcapi_mussad_svc_tenant.TradeActivityInfo)

	query := `SELECT TA.trade_activity_uno, trade_activity_name, trade_activity_code,
	TA.tenant_uno, TA.tenant_level_uno, TLL.tenant_level_name
	FROM tbl_mst_trade_activity TA
	INNER JOIN tbl_mst_trade_activity_locale TAL ON TA.trade_activity_uno=TAL.trade_activity_uno AND TAL.language_uno=@languageCode
	INNER JOIN tbl_mst_tenant_levels TL ON TA.tenant_level_uno=TL.tenant_level_uno AND TL.active=1
	INNER JOIN tbl_mst_tenant_levels_locale TLL ON TL.tenant_level_uno=TLL.tenant_level_uno AND TLL.language_uno=@languageCode
	WHERE TA.active=1 AND TL.tenant_uno=1 AND TL.level_uno=1 AND trade_activity_code = ANY(@tradeActivity)`

	args := pgx.NamedArgs{
		"tenantUno":     request.TenantUno,
		"languageCode":  request.LanguageCode,
		"tradeActivity": request.TradeActivity, // Pass the array directly
	}

	rows, err := sRepo.Client.DB.Query(ctx, query, args)

	log.Printf("3 rows: %v", rows)
	if err != nil {
		if err == pgx.ErrNoRows {
			return &grpcapi_mussad_svc_tenant.TradeActivityResponse{}, status.Errorf(codes.OK, "No data was found matching your request.")
		}
		slog.Error("Failed to execute query", slog.String("query", query), err.Error())
		return nil, status.Errorf(codes.Internal, "Failed to retrieve registration steps")
	}
	defer rows.Close()

	// Log each row
	for rows.Next() {
		var tradeActivityUno int32
		var tradeActivityName, tradeActivityCode, tenantLevelName string
		var tenantLevelUno int32

		err := rows.Scan(&tradeActivityUno, &tradeActivityCode, &tradeActivityName, &tenantLevelUno, &tenantLevelName)
		if err != nil {
			slog.Error("Failed to scan row", slog.String("error", err.Error()))
			return nil, status.Errorf(codes.Internal, "Failed to scan row")
		}

		// Process the row data as needed
		tradeActivityInfoMap[tradeActivityUno] = &grpcapi_mussad_svc_tenant.TradeActivityInfo{
			TradeActivityUno:  tradeActivityUno,
			TradeActivityName: tradeActivityName,
			TradeActivityCode: tradeActivityCode,
			TenantLevelUno:    tenantLevelUno,
			TenantLevelName:   tenantLevelName,
		}
	}

	

	if rows.Err() != nil {
		slog.Error("Error occurred during row iteration", slog.String("error", rows.Err().Error()))
		return nil, status.Errorf(codes.Internal, "Error occurred during row iteration")
	}
	// Populate the response
	resp.TradeActivityInfo = make([]*grpcapi_mussad_svc_tenant.TradeActivityInfo, 0, len(tradeActivityInfoMap))

	for _, step := range tradeActivityInfoMap {
		resp.TradeActivityInfo = append(resp.TradeActivityInfo, step)
	}

	return resp, nil
	//return resp, nil
}