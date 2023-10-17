package envc

import (
	envType "CSMS_ENV/env/type"
)

type EnvHttpRequestInterface interface {
	envType.ChargerInfoListReq |
		envType.ChargerInfoListAllReq |
		envType.ChargerStatusListReq |
		envType.ChargerStatusListAllReq |

		envType.ChargerInfoMyListReq |
		envType.ChargerStatusUpdateReq |

		envType.CardListReq |
		envType.CardListAllReq |
		envType.CardUpdateReq |
		envType.CardStatReq |

		envType.TradeListReq |
		envType.TradeListAllReq |
		envType.TradeRegisterReq |
		envType.TradeExlistReq |
		envType.TradeStatReq |

		envType.UseListReq |
		envType.UseRegisterReq |
		envType.UseDeleteReq |
		envType.UseStatReq
}

type EnvHttpResponseInterface interface {
	envType.ChargerInfoListRes |
		envType.ChargerInfoListAllRes |
		envType.ChargerStatusListRes |
		envType.ChargerStatusListAllRes |

		envType.ChargerInfoMyListRes |
		envType.ChargerStatusUpdateRes |

		envType.CardListRes |
		envType.CardListAllRes |
		envType.CardUpdateRes |
		envType.CardStatRes |

		envType.TradeListRes |
		envType.TradeListAllRes |
		envType.TradeRegisterRes |
		envType.TradeExlistRes |
		envType.TradeStatRes |

		envType.UseListRes |
		envType.UseRegisterRes |
		envType.UseDeleteRes |
		envType.UseStatRes
}
